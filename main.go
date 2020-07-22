package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)
	mux.HandleFunc("/pikachu", Pikachu)
	mux.HandleFunc("/bulbasaur", Bulbasaur)
	mux.HandleFunc("/squirtle", Squirtle)
	mux.HandleFunc("/charmander", Charmander)

	log.Println("Starting server :8080")

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,

		// setting custom timeouts
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

// resolveHosIP function
func resolveHostIP() string {

	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {

		networkIP, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !networkIP.IP.IsLoopback() && networkIP.IP.To4() != nil {

			ip := networkIP.IP.String()

			fmt.Println("Resolved Host IP: " + ip)

			return ip
		}
	}
	return ""
}

// Greet Function
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to POKEMON API from "+resolveHostIP())
}

// Pikachu Function
func Pikachu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "quu..__\n $$$b  `---.__\n  \"$$b        `--.                          ___.---uuudP\n   `$$b           `.__.------.__     __.---'      $$$$\"              .\n     \"$b          -'            `-.-'            $$$\"              .'|\n       \".                                       d$\"             _.'  |\n         `.   /                              ...\"             .'     |\n           `./                           ..::-'            _.'       |\n            /                         .:::-'            .-'         .'\n           :                          ::''\\          _.'            |\n          .' .-.             .-.           `.      .'               |\n          : /'$$|           .@\"$\\           `.   .'              _.-'\n         .'|$u$$|          |$$,$$|           |  <            _.-'\n         | `:$$:'          :$$$$$:           `.  `.       .-'\n         :                  `\"--'             |    `-.     \\\n        :##.       ==             .###.       `.      `.    `\\\n        |##:                      :###:        |        >     >\n        |#'     `..'`..'          `###'        x:      /     /\n         \\                                   xXX|     /    ./\n          \\                                xXXX'|    /   ./\n          /`-.                                  `.  /   /\n         :    `-  ...........,                   | /  .'\n         |         ``:::::::'       .            |<    `.\n         |             ```          |           x| \\ `.:``.\n         |                         .'    /'   xXX|  `:`M`M':.\n         |    |                    ;    /:' xXXX'|  -'MMMMM:'\n         `.  .'                   :    /:'       |-'MMMM.-'\n          |  |                   .'   /'        .'MMM.-'\n          `'`'                   :  ,'          |MMM<\n            |                     `'            |tbap\\\n             \\                                  :MM.-'\n              \\                 |              .''\n               \\.               `.            /\n                /     .:::::::.. :           /\n               |     .:::::::::::`.         /\n               |   .:::------------\\       /\n              /   .''               >::'  /\n              `',:                 :    .'\n                                   `:.:' Tim Park\n ")
}

// Bulbasaur Function
func Bulbasaur(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "                                           /\n                        _,.------....___,.' ',.-.\n                     ,-'          _,.--\"        |\n                   ,'         _.-'              .\n                  /   ,     ,'                   `\n                 .   /     /                     ``.\n                 |  |     .                       \\.\\\n       ____      |___._.  |       __               \\ `.\n     .'    `---\"\"       ``\"-.--\"'`  \\               .  \\\n    .  ,            __               `              |   .\n    `,'         ,-\"'  .               \\             |    L\n   ,'          '    _.'                -._          /    |\n  ,`-.    ,\".   `--'                      >.      ,'     |\n . .'\\'   `-'       __    ,  ,-.         /  `.__.-      ,'\n ||:, .           ,'  ;  /  / \\ `        `.    .      .'/\n j|:D  \\          `--'  ' ,'_  . .         `.__, \\   , /\n/ L:_  |                 .  \"' :_;                `.'.'\n.    \"\"'                  \"\"\"\"\"'                    V\n `.                                 .    `.   _,..  `\n   `,_   .    .                _,-'/    .. `,'   __  `\n    ) \\`._        ___....----\"'  ,'   .'  \\ |   '  \\  .\n   /   `. \"`-.--\"'         _,' ,'     `---' |    `./  |\n  .   _  `\"\"'--.._____..--\"   ,             '         |\n  | .\" `. `-.                /-.           /          ,\n  | `._.'    `,_            ;  /         ,'          .\n .'          /| `-.        . ,'         ,           ,\n '-.__ __ _,','    '`-..___;-...__   ,.'\\ ____.___.'\n `\"^--'..'   '-`-^-'\"--    `-^-'`.''\"\"\"\"\"`.,^.`.--' mh\n ")
}

// Squirtle Function
func Squirtle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "               _,........__\n            ,-'            \"`-.\n          ,'                   `-.\n        ,'                        \\\n      ,'                           .\n      .'\\               ,\"\".       `\n     ._.'|             / |  `       \\\n     |   |            `-.'  ||       `.\n     |   |            '-._,'||       | \\\n     .`.,'             `..,'.'       , |`-.\n     l                       .'`.  _/  |   `.\n     `-.._'-   ,          _ _'   -\" \\  .     `\n`.\"\"\"\"\"'-.`-...,---------','         `. `....__.\n.'        `\"-..___      __,'\\          \\  \\     \\\n\\_ .          |   `\"\"\"\"'    `.           . \\     \\\n  `.          |              `.          |  .     L\n    `.        |`--...________.'.        j   |     |\n      `._    .'      |          `.     .|   ,     |\n         `--,\\       .            `7\"\"' |  ,      |\n            ` `      `            /     |  |      |    _,-'\"\"\"`-.\n             \\ `.     .          /      |  '      |  ,'          `.\n              \\  v.__  .        '       .   \\    /| /              \\\n               \\/    `\"\"\\\"\"\"\"\"\"\"`.       \\   \\  /.''                |\n                `        .        `._ ___,j.  `/ .-       ,---.     |\n                ,`-.      \\         .\"     `.  |/        j     `    |\n               /    `.     \\       /         \\ /         |     /    j\n              |       `-.   7-.._ .          |\"          '         /\n              |          `./_    `|          |            .     _,'\n              `.           / `----|          |-............`---'\n                \\          \\      |          |\n               ,'           )     `.         |\n                7____,,..--'      /          |\n                                  `---.__,--.'mh\n")
}

// Charmander Function
func Charmander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "              _.--\"\"`-..\n            ,'          `.\n          ,'          __  `.\n         /|          \" __   \\\n        , |           / |.   .\n        |,'          !_.'|   |\n      ,'             '   |   |\n     /              |`--'|   |\n    |                `---'   |\n     .   ,                   |                       ,\".\n      ._     '           _'  |                    , ' \\ `\n  `.. `.`-...___,...---\"\"    |       __,.        ,`\"   L,|\n  |, `- .`._        _,-,.'   .  __.-'-. /        .   ,    \\\n-:..     `. `-..--_.,.<       `\"      / `.        `-/ |   .\n  `,         \"\"\"\"'     `.              ,'         |   |  ',,\n    `.      '            '            /          '    |'. |/\n      `.   |              \\       _,-'           |       ''\n        `._'               \\   '\"\\                .      |\n           |                '     \\                `._  ,'\n           |                 '     \\                 .'|\n           |                 .      \\                | |\n           |                 |       L              ,' |\n           `                 |       |             /   '\n            \\                |       |           ,'   /\n          ,' \\               |  _.._ ,-..___,..-'    ,'\n         /     .             .      `!             ,j'\n        /       `.          /        .           .'/\n       .          `.       /         |        _.'.'\n        `.          7`'---'          |------\"'_.'\n       _,.`,_     _'                ,''-----\"'\n   _,-_    '       `.     .'      ,\\\n   -\" /`.         _,'     | _  _  _.|\n    \"\"--'---\"\"\"\"\"'        `' '! |! /\n                            `\" \" -' mh\n")
}