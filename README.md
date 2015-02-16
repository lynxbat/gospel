Gospel

A library for collecting details about a system. Competely inspired by PuppetLabs' Facter(C) with the intention of providing a smiliar set of features for a Go language environemnt.

Core goals

* Fast execution
* CLI and Go API
* Extensible
* Simple

Like Facter has `facts` Gospel has `truths`.

Truths are bit of truth about the system you are collecting from in the form of a key/value pair.


WIP

Modules

* Gather - gathers truths and returns them through an Outputter
* Truth
  * packages that gather truths, buildtime separation by arch
  * runtime folder(s) for extenteding truths via YAML or JSON
* Speak - modules for printing, sending, saving truths
  * JSON
  * Terminal table
  * YAML