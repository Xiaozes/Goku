package Glogin

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"log"
	"time"
)

func Sshlogin(type_model, ssh_User, ssh_Password, ssh_Ip, cmd, ssh_Keypath, ssh_proxy string, ssh_Port int) {
	sshType := type_model
	sshUser := ssh_User
	sshPassword := ssh_Password
	sshIP := ssh_Ip
	sshPort := ssh_Port
	sshKeyPath := ssh_Keypath
	sshproxy := ssh_proxy
	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if sshType == "password" {
		log.Println("密码登录模式.......")
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else if sshType == "key" {
		log.Println("密钥登录模式.......")
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshKeyPath)}
	}
	addr := fmt.Sprintf("%s:%d", sshIP, sshPort)
	if sshproxy != "" {
		client, _ := proxiedSSHClient(ssh_proxy, addr, config)
		session, err := client.NewSession()
		if err != nil {
			log.Fatal("创建ssh session 失败", err)
			session.Close()
		}
		combo, err := session.CombinedOutput(cmd)
		if err != nil {
			log.Fatal("远程执行cmd 失败", err)
		}
		log.Println("命令输出:\n", string(combo))
	} else {
		sshClient, err := ssh.Dial("tcp", addr, config)
		if err != nil {
			log.Fatal("创建ssh client 失败", err)
			sshClient.Close()
		}
		defer sshClient.Close()
		session, err := sshClient.NewSession()
		if err != nil {
			log.Fatal("创建ssh session 失败", err)
			session.Close()
		}
		combo, err := session.CombinedOutput(cmd)
		if err != nil {
			log.Fatal("远程执行cmd 失败", err)
		}
		log.Println("命令输出:\n", string(combo))
	}

}

func proxiedSSHClient(proxyAddress, sshServerAddress string, sshConfig *ssh.ClientConfig) (*ssh.Client, error) {
	dialer, err := proxy.SOCKS5("tcp", proxyAddress, nil, proxy.Direct)
	if err != nil {
		log.Fatal("Proxy代理失败", err)
		return nil, err
	}

	conn, err := dialer.Dial("tcp", sshServerAddress)
	if err != nil {
		log.Fatal("创建ssh client 失败", err)
		return nil, err
	}

	c, chans, reqs, err := ssh.NewClientConn(conn, sshServerAddress, sshConfig)
	if err != nil {
		return nil, err
	}

	return ssh.NewClient(c, chans, reqs), nil
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}
