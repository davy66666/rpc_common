package k8sx

import (
	"fmt"

	"encoding/base64"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type KubeConf struct {
	Server      string // https://your-k8s-api-server
	BearerToken string // 服务账号Token
	CaCert      string // Base64或明文 CA 证书 PEM
}

func MustK8sClient(c KubeConf) *kubernetes.Clientset {
	// 构建 rest.Config
	cfg := &rest.Config{
		Host:        c.Server,
		BearerToken: c.BearerToken,
	}

	if c.CaCert != "" {
		// 解码 CA 证书（如果是 base64）
		caData := []byte(c.CaCert)
		if decoded, err := base64.StdEncoding.DecodeString(c.CaCert); err == nil {
			caData = decoded
		}

		cfg.TLSClientConfig = rest.TLSClientConfig{
			CAData: caData,
		}
	} else {
		cfg.TLSClientConfig.Insecure = true
	}

	// 创建 clientset
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to create k8s client: %w", err))
	}

	return clientset
}
