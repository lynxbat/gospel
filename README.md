##Gospel

Gospel is a library for collecting details about a system. It is competely inspired by [PuppetLabs' Facter(C)](https://github.com/puppetlabs/facter) with the intention of providing a similar set of features for a Go language environment.

###Core goals

* Fast execution
* CLI and Go API
* Extensible
* Simple

Like Facter has `facts`, Gospel has `truths`.

All `truths` are bits of truth about the system you are collecting from in the form of a key/value pair.


WIP

###Modules

* **Gather** - gathers truths and returns them through an Outputter
* **Truth**
  * packages that gather truths, buildtime separation by arch
  * runtime folder(s) for extenteding truths via YAML or JSON
* **Speak** - modules for printing, sending, saving truths
  * JSON
  * Terminal table
  * YAML

## Usage

WIP

## Contributing

WIP

## License and Authors

**Gospel** is Open Source software released under the Apache 2.0 License. Please see the [LICENSE](LICENSE) file for full license details.

* Author: Nicholas Weaver ([@lynxbat](http://github.com/lynxbat) | [email](lynxbat@gmail.com))