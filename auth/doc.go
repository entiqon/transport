/*
Package auth defines the authentication contract used by transport clients.

Authentication is applied by strategies that modify outgoing HTTP
requests before execution. Transport implementations use the Auth
interface to remain independent of specific authentication methods.
*/
package auth
