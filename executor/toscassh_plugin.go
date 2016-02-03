package executor

import (
	"bufio"
	"bytes"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/owulveryck/khoreia/choreography"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
)

// toscassh is a special plugin where all the inputs are passed as environment variables
func (n *node) toscassh() error {
	var conf map[string]sshConfig
	r, err := os.Open("sshConfig.yaml")
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return err
	}
	r.Close()

	if _, ok := conf[n.Target]; !ok {
		n.State = choreography.Failure
		return fmt.Errorf("Cannot find entry for host %v in the ssh config file", n.Target)
	}
	var auth []ssh.AuthMethod
	pubkey := PublicKeyFile(conf[n.Target].PrivateKeyFile)
	if pubkey != nil {
		auth = append(auth, pubkey)
	}
	agent := SSHAgent()
	if agent != nil {
		auth = append(auth, agent)
	}

	if len(auth) == 0 {
		n.State = choreography.Failure
		return fmt.Errorf("No authentication found for host %v", n.Target)
	}

	// ssh.Password("your_password")
	sshConfig := &ssh.ClientConfig{
		User: conf[n.Target].User,
		Auth: auth,
	}

	client := &SSHClient{
		Config: sshConfig,
		Host:   conf[n.Target].Host,
		Port:   conf[n.Target].Port,
	}

	command := ""
	for _, arg := range n.Args {
		command = fmt.Sprintf("%v ; echo \"%v\"", command, arg)
	}
	command = fmt.Sprintf("(%v ; cat %v && echo env) | /bin/ksh", command, n.Artifact)
	var outbuf bytes.Buffer
	//outbuf = *bytes.NewBuffer(output)
	cmd := &SSHCommand{
		Path:   command,
		Env:    []string{},
		Stdin:  os.Stdin,
		Stdout: &outbuf,
		Stderr: os.Stderr,
	}

	log.Printf("[%v] Running command: %s\n", n.Name, cmd.Path)
	n.State = choreography.Running
	if err := client.RunCommand(cmd); err != nil {
		n.State = choreography.Failure
		log.Printf("[%v] command run error: %s\n", n.Name, err)
		log.Println("Output:", outbuf.String())
		return err
	}
	// Now fill the output
	// find the variables
	for k, _ := range n.Outputs {
		out := outbuf
		re := regexp.MustCompile(fmt.Sprintf("^%v=(.*)", k))
		scanner := bufio.NewScanner(&out)
		for scanner.Scan() {
			txt := scanner.Text()
			//log.Printf("[%v] Output: %v", n.Name, txt)
			if re.MatchString(txt) {
				log.Printf("[%v] Output %v Matched argument %v", n.Name, txt, k)
				args := re.FindStringSubmatch(txt)
				n.Outputs[k] = args[1]
			}
		}
	}

	n.State = choreography.Success
	return nil
}
