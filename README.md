<div>
  <img alt="the speckle cube" src="https://user-images.githubusercontent.com/2679513/131189167-18ea5fe1-c578-47f6-9785-3748178e4312.png" width="150px"/><br/>
  Speckle | speckle-go 🐹
</div>

Speckle Go is a Golang SDK for interacting with Speckle, focusing on Speckle Automate functionality.

## Project Structure

- `cmd/examples`: Implementation examples
  - `speckle`: Client SDK usage
  - `speckle_automate`: Example Automate function implementation
- `internal`: Private application and library code
  - `api`: API-related code
  - `core`: Core functionality
  - `logging`: Logging utilities
  - `objects`: Object definitions
  - `serialization`: Serialization utilities
  - `transports`: Transport layer implementations
- `pkg`: Library code that's ok to use by external applications
  - `speckle`: Core SDK
  - `speckle_automate`: Automate function helpers
  - `objects`: Isolated Objects package
- `tests`: Additional external tests
- `go.mod`: 

### Key Components

- `pkg/speckle`: Overall Client SDK
- `pkg/speckle_automate`: The core functionality for creating and running Speckle Automate functions.
- `pkg/objects`: Defines Speckle object types and interfaces.

## Usage

To create a Speckle Automate function:

1. Define your function inputs:

```go
type FunctionInputs struct {
    Param1 string `json:"param1" validate:"required"`
    Param2 int    `json:"param2"`
}
```
2. Implement your function logic:
  ```go
  package main

  import (
      "github.com/specklesystems/speckle-go/pkg/automate"
  )
  
  func main() {
      automate.Run(automate.Function{
          Inputs: FunctionInputs{},
          Run: func(ctx *automate.Context[FunctionInputs]) error {
              // Your function logic here
              return nil
          },
      })
  }
```
--- 
<p ><b>Speckle</b> is the data infrastructure for the AEC industry.</p>

<p>
  <a href="https://twitter.com/SpeckleSystems"><img src="https://img.shields.io/twitter/follow/SpeckleSystems?style=social" alt="Twitter Follow"></a>
  <a href="https://speckle.community"><img src="https://img.shields.io/discourse/users?server=https%3A%2F%2Fspeckle.community&style=flat-square&logo=discourse&logoColor=white" alt="Community forum users"></a>
  <a href="https://speckle.systems"><img src="https://img.shields.io/badge/https://-speckle.systems-royalblue?style=flat-square" alt="website"></a>
  <a href="https://speckle.guide/dev/"><img src="https://img.shields.io/badge/docs-speckle.guide-orange?style=flat-square&logo=read-the-docs&logoColor=white" alt="docs"></a>
</p>

## About Speckle

What is Speckle? Check our![YouTube Video Views](https://img.shields.io/youtube/views/B9humiSpHzM?label=Speckle%20in%201%20minute%20video&style=social)

### Features

- **Object-based:** Say goodbye to files! Speckle is the first object-based platform for the AEC industry
- **Version control:** Speckle is the Git & Hub for geometry and BIM data
- **Collaboration:** Share your designs and collaborate with others
- **3D Viewer:** See your CAD and BIM models online, share and embed them anywhere
- **Interoperability:** get your CAD and BIM models into other software without exporting or importing
- **Real-time:** Get real-time updates and notifications and changes
- **GraphQL API:** Get what you need anywhere you want it
- **Webhooks:** The base for automation and next-gen pipelines
- **Built for developers:** We are building Speckle with developers in mind and have tools for every stack
- **Built for the AEC industry:** Speckle connectors are plugins for the most common software used in the industry, such as Revit, Rhino, Grasshopper, AutoCAD, Civil 3D, Excel, Unreal Engine, Unity, QGIS, Blender and more!

### Try Speckle now!

Give Speckle a try in no time by:

- [![speckle](https://img.shields.io/badge/https://-app.speckle.systems-0069ff?style=flat-square&logo=hackthebox&logoColor=white)](https://app.speckle.systems) ⇒ creating an account at our public server
- [![create a droplet](https://img.shields.io/badge/Create%20a%20Droplet-0069ff?style=flat-square&logo=digitalocean&logoColor=white)](https://marketplace.digitalocean.com/apps/speckle-server?refcode=947a2b5d7dc1) ⇒ deploying an instance in 1 click

### Resources

- [![Community forum users](https://img.shields.io/badge/community-forum-green?style=for-the-badge&logo=discourse&logoColor=white)](https://speckle.community) for help, feature requests or just to hang with other speckle enthusiasts, check out our community forum!
- [![website](https://img.shields.io/badge/tutorials-speckle.systems-royalblue?style=for-the-badge&logo=youtube)](https://speckle.systems) our tutorials portal is full of resources to get you started using Speckle
- [![docs](https://img.shields.io/badge/docs-speckle.guide-orange?style=for-the-badge&logo=read-the-docs&logoColor=white)](https://speckle.guide/dev/) reference on almost any end-user and developer functionality



### Style guide

We follow the standard Go style guidelines and use `gofmt` for code formatting. Run `gofmt` on your code before submitting a pull request.

## Contributing

Please make sure you read the [contribution guidelines](.github/CONTRIBUTING.md) and [code of conduct](.github/CODE_OF_CONDUCT.md) for an overview of the practices we try to follow.

## Community

The Speckle Community hangs out on [the forum](https://speckle.community). Join and introduce yourself, and feel free to ask us questions!

## Security

For security vulnerabilities or concerns, please contact us at security[at]speckle.systems.

## License

Unless otherwise described, the code in this repository is licensed under the Apache-2.0 License. Please note that some modules, extensions or code herein might be otherwise licensed. This is indicated in the folder's root under a different license file or the respective file's header. If you have questions, please contact us via [email](mailto:hello@speckle.systems).
