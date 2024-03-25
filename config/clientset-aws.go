package config

import (
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

func NewEKSClientset(options EKSOptions) (*kubernetes.Clientset, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(options.Region),
		Credentials: credentials.NewStaticCredentials(options.AccessKey, options.SecretKey, ""),
	}))
	eksSvc := eks.New(sess)

	input := &eks.DescribeClusterInput{
		Name: aws.String(options.ClusterName),
	}

	result, err := eksSvc.DescribeCluster(input)
	if err != nil {
		return nil, fmt.Errorf("error describing cluster: %s", err)
	}

	gen, err := token.NewGenerator(true, false)
	if err != nil {
		return nil, err
	}

	opts := &token.GetTokenOptions{
		ClusterID: aws.StringValue(result.Cluster.Name),
		Session:   sess,
	}
	tok, err := gen.GetWithOptions(opts)
	if err != nil {
		return nil, err
	}

	ca, err := base64.StdEncoding.DecodeString(aws.StringValue(result.Cluster.CertificateAuthority.Data))
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(
		&rest.Config{
			Host:        aws.StringValue(result.Cluster.Endpoint),
			BearerToken: tok.Token,
			TLSClientConfig: rest.TLSClientConfig{
				CAData: ca,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
