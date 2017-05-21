# imagetoascii - a service to convert images to ASCII "art"

## Build

You can build with make:

    (Set GOPATH if not already set. $HOME is a good spot)
    $ export GOPATH=/your/gopath

Now you can build with:

    $ make clean
    $ make

The binary is now in the build directory, which you can run with:

    $ ./build/imagetoascii

It serves on port 8080 by default

See --help for more usage information:

    $ ./build/imagetoascii --help

## Usage

This service exposes a single API endpoint /toascii that can be POSTed to. It
expects a JSON payload with a single data element that contains a base64 encoded
image. The image can be a jpeg, png, or gif. An example:

    {"data": "R0lGODdhFQAbAJEAAAAAAP///wAAAAAAACH5BAkAAAIALAAAAAAVABsAAAJVjI+pC+29EjAT1VjzvcfJAFEW5olUpoQmGGEdmY6dKs+l+7Un88A6lyMBV7HiEERb0UJAjpOlrO0mHgiV2KNeta1NtqtqSnqn24a1RUO9j7b7qEsUAAA7"}

So if you would like to call the service with curl, you could do:

    $ curl -s -X POST localhost:8080/toascii -H "Content-Type: application/json" --data '{"data": "R0lGODdhFQAbAJEAAAAAAP///wAAAAAAACH5BAkAAAIALAAAAAAVABsAAAJVjI+pC+29EjAT1VjzvcfJAFEW5olUpoQmGGEdmY6dKs+l+7Un88A6lyMBV7HiEERb0UJAjpOlrO0mHgiV2KNeta1NtqtqSnqn24a1RUO9j7b7qEsUAAA7"}'
    {"data":".......nnnnnnnn......\n.....nn.nn.nn..nn....\n.nnnn.....$n$$...$...\n$n..n......nn....$...\nnn....nnnnn.....nn...\n.n$nnn.........n$....\n...11111011.....110..\n.....nn$......nn..$n.\n..nn$..nn...nn...nn..\n..n$n$nn$nn$nnnnnnnn.\n......n$n..n$nnn.nn..\n..nn$n$nn.nn$n$nn..$.\n..nnnnnnnnnnnnnnnnnn.\n"}

Not very satisfying. But if you install jq, and do some clever command line trickery, you can feed any image into the service and parse the json:


    $ curl -s -X POST localhost:8080/toascii -H "Content-Type: application/json" --data "{\"data\": \"$(base64 agumon.gif)\"}" | jq .data -r
    ..........................................
    ..............$$$$$$$$$$$$$$$$............
    ............$$....$$..........$$..........
    ..........$$....$$....$$$$......$$........
    ....$$$$$$..........$$..$$$$......$$......
    ..$$................$$$$$$$$......$$......
    $$$$....$$............$$$$........$$......
    $$................................$$......
    $$................$$$$............$$......
    ..$$........$$$$$$..............$$........
    ..$$$$$$$$$$....................$$........
    ....$$........................$$$$........
    ......$$$$$$$$$$................$$$$......
    ................$$..................$$....
    ................$$$$$$..............$$....
    ..............$$..............$$....$$....
    ..........$$$$$$............$$......$$$$..
    ......$$$$....$$..........$$........$$....
    ....$$..$$......$$......$$........$$......
    ....$$$$..$$..$$$$$$..$$$$..$$..$$....$$..
    ......$$$$$$$$..$$..$$$$..$$..$$..$$$$....
    ..............$$......$$$$$$$$$$..$$......
    ............$$$$$$......$$..........$$....
    ......$$$$$$$$$$......$$$$$$$$$$......$$..
    ....$$..$$..$$..$$..$$..$$..$$..$$....$$..
    ....$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$..
    ..........................................


Images greater than 80 pixels wide will be resized to fit in 80 ASCII columns.

## Test

You can run tests with:

    $ make test

## Dependencies

imagetoascii has two main non-stdlib dependencies: echo and nfnt/resize

* echo is a simple web framework for Go. It is an active project with many
  questions on Stack Overflow and lots of activity on github (7000+ stars and
  has accepted almost 300 PRs as of this writing). It is used to handle and
  route HTTP requests.

* nfnt/resize is a simple image resize library. It is a popular project,
  however it doesn't have many recent commits. That said it just provides two
  simple resize functions, so it may be that it simply doesn't have many bugs
  or need any new features.

## Building for Docker

To build a Docker image, simply:

    $ make docker

And you can run it with:

    $ docker run -p 8080:8080 --rm --name imagetoascii imagetoascii

You can pass arguments like:

    $ docker run -p 8000:8000 --name imagetoascii --rm imagetoascii --version
