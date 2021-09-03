# PROJECT_FINAL_GROUP03

[![CodeQL](https://github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/actions/workflows/codeql-analysis.yml/badge.svg?branch=main)](https://github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/actions/workflows/codeql-analysis.yml)

[![Lint, Build & Tests - Go](https://github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/actions/workflows/go-ci.yml/badge.svg)](https://github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/actions/workflows/go-ci.yml)

## Setup

### Invite bot

To invite the discord bot to your server, follow [this link](https://discord.com/api/oauth2/authorize?client_id=882205244170334218&permissions=54177037431&scope=bot), choose your server in the drop down menu, and press the validation button. 
You can then choose the bot permissions - if you want to change them - before clicking authorize.

### Run project locally

the following make commands are available :

|command|description|
|-|-|
| **make start**| start the project with docker compose |
| **make stop**| stop the project |
| **make log**| follow log output |
| **make init**| same as start, but also copy files |
| **make copy-files**| copy env files to the right directories, concat global env to other env files |

## Bot commands

Once the bot is in the discord server, you can address a message to the bot by beginning your message with `/admin`

<details>
  <summary><b>Login command</b></summary>

  ```
    /admin login
  ```

  You will receive a private message containing a link to the front-end interface, allowing you to manage your server.
</details>

## Contributing

We follow a [code of conduct](CODE_OF_CONDUCT.md), if you wish to contribute on this project, we strongly advise you to read it.

<details>	
  <summary><b>Branch naming convention</b></summary>

- You branch should have a name that reflects it's purpose.

- It should use the same guidelines as [COMMIT_CONVENTIONS](COMMIT_CONVENTIONS.md) (`feat`, `fix`, `build`, `perf`, `docs`), followed by an underscore (`_`) and a very quick summary of the subject in [kebab case][1].

    Example: `feat_add-image-tag-database-relation`.
</details>
<details>
  <summary><b>Pull requests and commits</b></summary>

Pull requests in this project follow two conventions, you will need to use the templates available in the [ISSUE_TEMPLATE](.github/ISSUE_TEMPLATE) folder :

- Adding a new feature should use the [FEATURE_REQUEST](.github/ISSUE_TEMPLATE/feature_request.md) template.
- Reporting a bug should use the [BUG_REPORT](.github/ISSUE_TEMPLATE/bug_report.md) template.

If your pull request is still work in progress, please add "WIP: " (Work In Progress) in front of the title, therefor you inform the maintainers that your work is not done, and we can't merge it.

The naming of the PR should follow the same rules as the [COMMIT_CONVENTIONS](COMMIT_CONVENTIONS.md)
</details>

## Continuous Integration (CI)

A CI pipeline is configured for this project and is accessible in the [Go-CI](.github/workflows/go-ci.yaml) file.

The pipeline will run 4 different jobs:

- Dependencies check
- Linter
- Build
- Tests

The pipeline will be triggered automatically when creating a new **Pull Request** and on each **push** on it. It will also be triggered on push on `main` branch.

## Contributors

<table align="center">
  <tr>
    <td align="center">
    <a href="https://github.com/jasongauvin">
      <img src="https://avatars1.githubusercontent.com/u/41618366?s=400&u=b970ed03cbb921ce1312ef86b39093e4fa0be7e3&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Jason Gauvin</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/JackMaarek/">
      <img src="https://avatars3.githubusercontent.com/u/28316928?s=400&u=3cdfb5b0683245ad333a39cfca3a5251f3829824&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Jacques Maarek</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/edwinvautier">
      <img src="https://avatars3.githubusercontent.com/u/35581502?s=460&u=d9096f90151f35552d9adcd57bacaee366f0aaef&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Edwin Vautier</b></sub>
    </a>
    </td>
    </td>
        <td align="center">
        <a href="https://github.com/AlexandreLch">
          <img src="https://avatars.githubusercontent.com/u/25430432?v=4" width="100px;" alt=""/>
          <br />
          <sub><b>Alexandre Lellouche</b></sub>
        </a>
        </td>
  </tr>
</table>
