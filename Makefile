build:
	go build

demo: build
	Xephyr -br -ac -noreset -screen 960x540 :1 &
	sleep 1
	DISPLAY=:1 ./.xinitrc
