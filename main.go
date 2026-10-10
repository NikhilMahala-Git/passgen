package main

import (
	"fmt"
	"os"
	"strings"
)

const PASSWORD_CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_";
func passgen(input string) string {
    var _password1 = "";
    var _password2 = "";

    var _sum = 0;
	for i := 0; i < len(input); i++ {
        _sum += int(input[i]);
		var _index int = strings.Index(PASSWORD_CHARS, input[i:i+1])
        if _index == -1 {
            _index = 0;
        }
        var _newIndex = _index + (_sum % (len(PASSWORD_CHARS) - _index));
		_password2 += PASSWORD_CHARS[_newIndex:_newIndex + 1];
    }

	for i := 0; i < len(input); i++ {
		var _index int = strings.Index(PASSWORD_CHARS, input[i:i+1])
        if _index == -1 {
            _index = 0;
        }
        var _newIndex = _index + (_sum % (len(PASSWORD_CHARS) - _index));
		_password1 += PASSWORD_CHARS[_newIndex:_newIndex + 1];
    }

    return (_password1 + _password2);
}


func main() {
	_args := os.Args;
	_argCount := len(_args)

	if _argCount < 3 {
		fmt.Printf( "Usage: passgen <prefix> <password>\n");
		return
	}

	_prefix := _args[1]
	_passwd := _args[2]

	fmt.Printf("Generated: %v%v\n", _prefix, passgen(_passwd))
}
