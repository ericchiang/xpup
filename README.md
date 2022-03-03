# xpup

xpup is pup for XML.

It uses XPath rather than CSS selectors.

## Install

```
go get github.com/ericchiang/xpup
```

Binary installs coming soon.

## Example

```
$ curl -s https://www.w3schools.com/xml/note.xml
<?xml version="1.0" encoding="UTF-8"?>
<note>
  <to>Tove</to>
  <from>Jani</from>
  <heading>Reminder</heading>
  <body>Don't forget me this weekend!</body>
</note>

$ curl -sL https://www.w3schools.com/xml/note.xml | xpup '/*/body'
Don't forget me this weekend!
```
