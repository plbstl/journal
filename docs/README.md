<p align="center">
 <img src="https://upload.wikimedia.org/wikipedia/commons/6/63/Journal_Icon.svg" height=200 alt="Journal"/>
</p>

<h1 align="center">Journal</h1>
<p align="center">📔 Preserve your memories, self-reflect and record ideas on-the-go.</p>

<div align="center">

[![Release][release-badge]][release-link]
[![GitHub Issues][gh-issues-badge]][gh-issues-link]
[![GitHub Pull Requests][gh-pulls-badge]][gh-pulls-link]
[![GitHub contributors][gh-contributors-badge]][gh-contributors-link]
[![License Badge][license-badge]][license-link]
[![PkgGoDev][pkg-godev-badge]][pkg-godev-link]

</div>

---

Journal helps you manage multiple journals with ease from the comfort of your terminal, web browser or API client.
You can import/export journals as [horcruxes](https://en.wikipedia.org/wiki/Magical_objects_in_Harry_Potter#Horcruxes) and set simple customizations for layout, theme, and keybindings.

Journal brings out the quintessence of your potent yet pervious mind.
Get personal with it and use its gorgeous powers for good.

<!-- @todo: create screencasts for cli & web -->

<div align="center">

Cli Screencast

<!-- ![Cli Screencast](/resources/cli-demo.gif "-gifcontrol-mode=click;") -->

![Cli Screencast](https://media.giphy.com/media/jRlP4zbERYW5HoCLvX/giphy.gif '-gifcontrol-mode=click;')

Web Screencast

<!-- ![Web Screencast](/resources/web-demo.gif "-gifcontrol-mode=click;") -->

![Web Screencast](https://media.giphy.com/media/IwTWTsUzmIicM/giphy.gif '-gifcontrol-mode=click;')

</div>

## Quick Start <a name="quick-start"></a>

<!-- @todo: upload live demo; edit live-demo-link -->

Check out [live demo][live-demo-link] of web app.

### Installing

#### Binary Release

You can download a binary release for Windows, Mac or Linux [here][release-link].

#### Using Go Get

```sh
go get -u github.com/paulebose/journal
```

If you get an error claiming that _journal_ command cannot be found or is not defined, you
may need to add `~/go/bin` to your _\$PATH_. `~/go/bin` should not be mistaken for `$GOROOT/bin` (which is for Go's own binaries,
not installed apps like _journal_).

## Usage <a name="usage"></a>

Here is some basic usage information for terminal.

### Commands

Create and open a new note with the default author:

```sh
journal new
```

Create a new note with a custom id:

```sh
journal create note --id custom-id
```

Create a new author:

```sh
journal create author
```

Edit an author or tag:

```sh
journal edit <id>
```

Run `journal help` for more information.

```sh
journal help
```

### Keybindings

<pre>

  <kbd>r</kbd> refresh list
  <kbd>h</kbd> show help
  <kbd>t</kbd> toggle screen mode (normal/half/fullscreen)
  <kbd>,</kbd> show settings
  <kbd>/</kbd> search
  <kbd>x</kbd> select item
  <kbd>Q</kbd> quit
  <kbd>esc</kbd> close current panel

</pre>

Check the [keybindings docs](/keybindings.md) for more information.

### Configuration

Default path for config file:

-   Windows: `C:\Users\USERNAME\AppData\Roaming\Journal\config.yml`
-   MacOS: `/Users/USERNAME/Library/Application Support/Journal/config.yml`
-   Linux: `/home/USERNAME/.config/Journal/config.yml`

```yaml
# Example Config
cli:
    mouseEnabled: true
    editor: 'in-built'
    confirmOnQuit: false
    keybindings:
        # command: 'key'
web:
    port: 9001
```

Check the [configuration docs](/config.md) for more information.

[license-badge]: https://img.shields.io/github/license/paulebose/journal
[license-link]: LICENSE.md
[live-demo-link]: https://example.org
[pkg-godev-badge]: https://pkg.go.dev/badge/github.com/PaulEbose/journal
[pkg-godev-link]: https://pkg.go.dev/github.com/PaulEbose/journal
[gh-contributors-badge]: https://img.shields.io/github/contributors/paulebose/journal
[gh-contributors-link]: https://github.com/PaulEbose/journal/graphs/contributors
[gh-issues-badge]: https://img.shields.io/github/issues/paulebose/journal
[gh-issues-link]: https://github.com/paulebose/journal/issues
[gh-pulls-badge]: https://img.shields.io/github/issues-pr/paulebose/journal
[gh-pulls-link]: https://github.com/paulebose/journal/pulls
[release-badge]: https://img.shields.io/github/v/release/PaulEbose/journal?sort=semver
[release-link]: https://github.com/PaulEbose/journal/releases
