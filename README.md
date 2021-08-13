<p align="center">
 <img src="https://upload.wikimedia.org/wikipedia/commons/6/63/Journal_Icon.svg" height=200 alt="Journal"/>
</p>

<h3 align="center">Journal</h3>

<div align="center">

[![Release][release-badge]][release-link]
[![GitHub Issues][gh-issues-badge]][gh-issues-link]
[![GitHub Pull Requests][gh-pulls-badge]][gh-pulls-link]
[![GitHub contributors][gh-contributors-badge]][gh-contributors-link]
[![License Badge][license-badge]][license-link]
[![PkgGoDev][pkg-godev-badge]][pkg-godev-link]

</div>

---

<p align="center">📔 Preserve your memories, self-reflect and record ideas on-the-go.</p>

## 📝 Table of Contents

-   [About <a name = "about"></a>](#about)
-   [Getting Started](#getting-started)
-   [Usage](#usage)
-   [Contributing](#contributing)
-   [Authors](#authors)
-   [Acknowledgement](#acknowledgement)

## 🧐 About <a name = "about"></a>

Journal helps you manage multiple journals with ease from the comfort of your terminal, web browser or API client.
You can import/export journals as [horcruxes](https://en.wikipedia.org/wiki/Magical_objects_in_Harry_Potter#Horcruxes) and set simple customizations for layout, theme, and keybindings.

Journal brings out the quintessence of your potent yet pervious mind.
Get personal with it and use its gorgeous powers for good.

<!-- @todo: create screencasts for cli & web -->

<div align="center">

Cli Screencast

<!-- ![Cli Screencast](docs/resources/cli-demo.gif) -->

![Cli Screencast](https://media.giphy.com/media/11PEptfDmR4vjW/giphy.gif)

Web Screencast

<!-- ![Web Screencast](docs/resources/web-demo.gif) -->

![Web Screencast](https://media.giphy.com/media/IwTWTsUzmIicM/giphy.gif)

</div>

## 🏁 Getting Started <a name="getting-started"></a>

<!-- @todo: upload live demo; edit live-demo-link -->

Check out [live demo][live-demo-link] of web app.

### Installing

#### Binary Release

You can download a binary release for Windows, Mac or Linux [here]([release-link]).

#### Using Go Get

```sh
go get -u github.com/paulebose/journal
```

If you get an error claiming that _journal_ command cannot be found or is not defined, you
may need to add `~/go/bin` to your _\$PATH_. `~/go/bin` should not be mistaken for `$GOROOT/bin` (which is for Go's own binaries,
not installed apps like _journal_).

## 🎈 Usage <a name="usage"></a>

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

Check the [keybindings docs](docs/KEYBINDINGS.md) for more information.

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

Check the [configuration docs](docs/CONFIG.md) for more information.

## 🤿 Contributing <a name = "contributing"></a>

Contributors are welcome! Please check out the [contributing guide](CONTRIBUTING.md).
For discussions about stuff, you can use [GitHub Discussions][gh-discussions-link]

## ✍️ Authors <a name = "authors"></a>

-   [@PaulEbose](https://github.com/PaulEbose) - Idea & initial work

See also the list of [contributors][gh-contributors-link] who participated in this project.

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

-   Influence: [@Richmond-Eribo](https://github.com/Richmond-Eribo)

[license-badge]: https://img.shields.io/github/license/paulebose/journal
[license-link]: LICENSE
[live-demo-link]: https://example.org
[pkg-godev-badge]: https://pkg.go.dev/badge/github.com/PaulEbose/journal
[pkg-godev-link]: https://pkg.go.dev/github.com/PaulEbose/journal
[gh-contributors-badge]: https://img.shields.io/github/contributors/paulebose/journal
[gh-contributors-link]: https://github.com/PaulEbose/journal/graphs/contributors
[gh-discussions-link]: https://github.com/PaulEbose/journal/discussions
[gh-issues-badge]: https://img.shields.io/github/issues/paulebose/journal
[gh-issues-link]: https://github.com/paulebose/journal/issues
[gh-pulls-badge]: https://img.shields.io/github/issues-pr/paulebose/journal
[gh-pulls-link]: https://github.com/paulebose/journal/pulls
[release-badge]: https://img.shields.io/github/v/release/PaulEbose/journal?sort=semver
[release-link]: https://github.com/PaulEbose/journal/releases
