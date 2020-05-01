# k8s-enhancements
A tool to help track and manage [kubernetes/enhancements](https://github.com/kubernetes/enhancements) repo and tracking sheets.

# Setup

## Pre-requisites
1. Access to GitHub
2. Access to [kubernetes/ehnacements tracking sheet](https://docs.google.com/spreadsheets/d/1E89GNZRwmnfVerOqWqrmzVII8TQlQ0iiI_FgYKxanGs)
3. API Key for Accessing Google Sheets
4. OAuth2.0 Credentials for Google Developer Console for an APP that enables Spreadsheets API

## Setup
1. Create the OAuth2.0 Credentials in `$HOME/.k8s-enhancements/credentials.json`
2. Create `$HOME/.k8s-enhancements/config.yaml` and update Google Sheets API Key and GitHub Access Token
    ```yaml
    api-key: xyzk1
    git-access-token: xyzk1
    ```

## Templates
This tool also provides an easy way to interact with the Issue owners/assignees in the form of a template based comment generator for GitHub Issues. 

Once you review the KEP and Issue details manually, you can use one of the custom templates to initiate a message on the GitHub Issue. Templates can be put under `~/.k8s-enhancements/templates/` path

The name of the template file is the name you can use while invoking the command. Every template is rendered with a common `Mentions` parameter which is a
list of GitHub users `@` prefixed to the name. Multiple names are split with `/` 

### Example Template
Save this file as `~/.k8s-enhancements/templates/initial`. Now, you can use the `--template initial` as an argument to invoke the CLI.

```markdown
Hey there {{.Mentions}} -- 1.19 Enhancements shadow here. I wanted to check in and see if you think this Enhancement will be graduating in 1.19?

In order to have this part of the release:
1. The KEP PR must be merged in an `implementable` state
2. The KEP must have test plans
3. The KEP must have graduation criteria.

The current release schedule is:
- Monday, April 13: Week 1 - Release cycle begins
- Tuesday, May 19: Week 6 - Enhancements Freeze
- Thursday, June 25: Week 11 - Code Freeze
- Thursday, July 9: Week 14 - Docs must be completed and reviewed
- Tuesday, August 4: Week 17 - Kubernetes v1.19.0 released


If you do, I'll add it to the 1.19 tracking sheet (http://bit.ly/k8s-1-19-enhancements). Once coding begins please list all relevant k/k PRs in this issue so they can be tracked properly. 👍

Thanks!
```

### Usage
```bash
 ▲ github.com/harshanarayana/k8s-enhancements ✱❄︎ ./k8s-enhancements git issues comment --git-issue 1611 --mention jayunit100 --template initial                                                        ⇡ master :: 12h :: ⬢
```