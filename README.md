# Dreamers

Source for sample website build by Echo (Golang framework)

### Description

The concept is simple, I create a big night sky then people can send their stories to the stars. Everytime page reload,
the position of the stars will be changed randomly. The star that contain the story can be clicked to view. It's totally random so your stories are safe with the stars.

- Click on the sky
- Write your story then send
- Booom! Your story have been placed among the stars

### Installation

Clone the project:

```
git clone https://github.com/M15t/dreamers.git
```

Update 2018-09-23

> If you clone this project on a PC that run Windows OS. When you run "go run main.go", in my case, it'll display an error "exec: "gcc": executable file not found in %PATH%". So the fix is install http://tdm-gcc.tdragon.net/download.
> This error cause I use go-sqlite3 so I need a C compiler to use this driver.

### Run

Move to folder:

```
$ cd dreamers
```

Run main.go

```
$ go run main.go
```

Then visit the link http://localhost:9013

### References

https://github.com/labstack/echo

https://github.com/foolin/echo-template

https://github.com/jinzhu/gorm

https://github.com/asaskevich/govalidator
