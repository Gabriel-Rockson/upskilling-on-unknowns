package main

import "fmt"

type authenticationInfo struct {
  username string
  password string
}

func (authInfo authenticationInfo) getBasicAuth() string  {
  return fmt.Sprintf("Authorization: Basic %s:%s", authInfo.username, authInfo.password)
}

func main()  {
  robotUser := authenticationInfo{
    username: "Angelina",
    password: "HoolaaHoolaa",
  }

  robotUserBasicAuth := robotUser.getBasicAuth()

  fmt.Println(robotUserBasicAuth)
}
