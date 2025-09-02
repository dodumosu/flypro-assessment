package repositories

import "fmt"

var ErrEntityExists = fmt.Errorf("the specified entity already exists")
var ErrEntityDoesNotExist = fmt.Errorf("the specified entity does not exist")
