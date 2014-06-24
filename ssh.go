/**
 * asink v0.0.2-dev
 *
 * (c) Ground Six
 *
 * @package asink
 * @version 0.0.2-dev
 *
 * @author Harry Lawrence <http://github.com/hazbo>
 *
 * License: MIT
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
*/

package main

import (
    "fmt"
	"log"
    "io/ioutil"
    "bufio"
    "time"
	"./vendor/go.crypto/ssh"
)

/**
 * @var String remote name
 * @var String remote host
 * @var String remote port number
 * @var String remote username
 * @var String remote password
 */
type Remote struct {
    Name     string
    Host     string
    Port     string
    User     string
    Password string
    Key      ssh.Signer
}

var remotes  map[string]*Remote = nil
var sessions map[string]*ssh.Session = nil

/**
 * Inits remotes and sessions then
 * returns a new instance of remote
 *
 * @return *Remote new remote
 */
func NewRemote() *Remote {
    remotes  = make(map[string]*Remote)
    sessions = make(map[string]*ssh.Session)
    return new(Remote)
}

/**
 * Adds a new remote to the map of
 * remotes with a string key
 *
 * @param String remote name
 * @param String remote host
 * @param String remote port number
 * @param String remote username
 * @param String remote password
 *
 * @return nil
 */
func (r *Remote) AddRemote(name string, host string, port string, user string, password string) {
    remote := new(Remote)

    remote.Name     = name
    remote.Host     = host
    remote.Port     = port
    remote.User     = user
    remote.Password = password

    remotes[name] = remote
}

/**
 * Parses then adds the key to our remote struct
 *
 * @param String remote name
 * @param String key path
 *
 * @return nil
 */
func (r *Remote) AddSshKey(name string, file string) {
    remote    := remotes[name]
    remote.Key = parseKey(file)
}

/**
 * Starts a new SSH session which is then
 * stored in the sessions map with a given
 * key
 *
 * @param String session key
 *
 * @return nil
 */
func StartSession(name string) {
    remote := remotes[name]

    config := &ssh.ClientConfig{
        User: remote.User,
        Auth: []ssh.AuthMethod{
            ssh.Password(remote.Password),
            ssh.PublicKeys(remote.Key),
        },
    }

    hostname := remote.Host + ":" + remote.Port

    conn, err := ssh.Dial("tcp", hostname, config)
    if err != nil {
        log.Fatalf("unable to connect: %s", err)
    }

   //defer conn.Close()

    session, err := conn.NewSession()
    if err != nil {
        log.Fatalf("unable to create session: %s", err)
    }

    //defer session.Close()

    modes := ssh.TerminalModes{
        ssh.ECHO:          0,     // disable echoing
        ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
        ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
    }

    if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
        log.Fatalf("request for pseudo terminal failed: %s", err)
    }

    sessions[name] = session
}

/**
 * Runs the remote command given the session
 * key
 *
 * @param String session name
 * @param String command string
 *
 * @return nil
 */
func RunRemoteCommand(name string, command string) {
    session  := sessions[name]

    n, _ := session.StdoutPipe()
    err  := session.Start(command);

    if (err != nil) {
        panic(err)
    }
    for {
        r := bufio.NewReader(n)
        i, _, _ := r.ReadLine()

        result := string(i)
        if (result != "") {
            fmt.Println(string(i))
            time.Sleep(time.Millisecond * 50)
        } else {
            break
        }
    }
    session.Wait()
}

/**
 * Parses the key for the client so
 * we can SSH into the remote
 *
 * @param String file path
 *
 * @return ssh.Signer
 */
func parseKey(file string) ssh.Signer {
    privateBytes, err := ioutil.ReadFile(file)
    if err != nil {
        panic("Failed to load private key")
    }

    private, err := ssh.ParsePrivateKey(privateBytes)
    if err != nil {
        panic("Failed to parse private key")
    }
    return private
}
