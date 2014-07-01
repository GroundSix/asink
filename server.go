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
	"net/http"
	"io/ioutil"
	"./vendor/mux"
)

/**
 * Starts a very basic http server to
 * accept JSON input instead of a
 * static configuration file
 *
 * @return nil
 */
func startServer() {
    r := mux.NewRouter()
    r.HandleFunc("/", FetchJsonBody)
    http.Handle("/", r)
    http.ListenAndServe(":9000", nil)
}

/**
 * Fetches the body sent in the http
 * request and returns it as a string
 *
 * @param http.ResponseWriter
 * @param *http.Request
 */
func FetchJsonBody(w http.ResponseWriter, r *http.Request) {
	request_body, _ := ioutil.ReadAll(r.Body)
	initAsinkWithServer(string(request_body))
}
