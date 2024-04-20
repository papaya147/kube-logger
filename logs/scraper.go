package logs

import (
	"bytes"
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Scrape() error {
	for _, pod := range pods {
		go scrape(context.Background(), pod)
	}

	forever := make(chan bool)
	<-forever

	return nil
}

func scrape(ctx context.Context, pod NamespacedPod) error {
	req := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &corev1.PodLogOptions{
		SinceTime: &metav1.Time{Time: time.Now()},
		Follow:    true,
	})
	readCloser, err := req.Stream(ctx)
	if err != nil {
		return err
	}
	defer readCloser.Close()

	buf := make([]byte, 10240)
	for {
		_, err = readCloser.Read(buf)
		if err != nil {
			return err
		}

		n := bytes.IndexByte(buf[:], 0)

		if n <= 0 {
			continue
		}

		if err := write(pod.Namespace, pod.Name, buf[:n]); err != nil {
			return err
		}

		buf = make([]byte, 10240)
	}
}
