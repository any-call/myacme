package myacme

import (
	"crypto/tls"
	"golang.org/x/crypto/acme/autocert"
)

// 自动证书管理函数
func AutoCertManager(saveCertDir, certDomain string) *tls.Config {
	// 自动申请并管理证书
	certManager := &autocert.Manager{
		Cache:      autocert.DirCache(saveCertDir),     // 证书存储目录
		Prompt:     autocert.AcceptTOS,                 // 自动接受 Let's Encrypt 条款
		HostPolicy: autocert.HostWhitelist(certDomain), // 允许的域名
	}

	return &tls.Config{GetCertificate: certManager.GetCertificate}
}
