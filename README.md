# Amp Template Development Environment

this is a development environment for amp templates. 

## Setup
* Clone this repo to your computer

* download [Golang](https://golang.org) for your operating system

* download [Node LTS](https://nodejs.org/en/) for your operating system

* install yarn
    ```
    > npm install yarn -g
    ```

## Building Frontend Dependencies
* Open a terminal window and navigate to the root of the project
* Navigate to the frontend folder:
    ```
    > cd frontend    
    ```
* Run yarn to install dependencies:
    ```
    > yarn
    ```
* Run yarn in watch mode while developing:
    ```
    > yarn watch
    ```

## Frontend Structure
We use webpack to build scss/sass files in the `frontend/src/` folder and combine them into a `main.css` file in the `frontend/dist` directory.

The main file that the build starts with is `frontend/src/index.scss` so include any extra sass from there. 

You can create your own folder structure under `frontend/src` and just include it in `frontend/src/index.scss`.

## Template Structure

This web server uses [golang's templating engine](https://golang.org/pkg/text/template/). You can read up on it, but you're really unlikely to need it.
Just note that template syntax is `{{FunctionName}}`. 

For instance `{{GetAmpStyles}}` in the `templates/_head.html` file gets the built main.css file's contents and dumps it into the template. 

### NOTE!
Hot reloading is enabled by default as long as the server is running in development mode.

## Starting up the Server

* From the root directory of the project run:
    ```
    > go run server/main.go
    ```
  
  That should be it! you should see something like this:
    ```
    WARNING  02/15/20 5:50:25 PM [amp-templates Server listening at localhost:8080]
    ```