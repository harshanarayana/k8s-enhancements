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
