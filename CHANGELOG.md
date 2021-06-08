# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2021-06-01
### Changed
- golang version to 1.16
- golangspell-core templates
- github actions
- change from nlopes/slack to slack-go/slack
- change tests to assert lib
- change telegram webhook struct to receive more data
- fix execute and silence verb
- in gateway directory split between integrations (monitoring softwares) and chat

### Added
- google chat integration
- alert manager integration
- user and password authentication option for sensu api
- test directory with mock interfaces
- new config variables for security options: 
    - SENSUBOT_BLOCKED_VERBS: blocked list of verbs (get, execute, silence, delete, resolve)
    - SENSUBOT_BLOCKED_RESOURCES: blocked list of resources from sensu api
    - SENSUBOT_SLACK_ADMIN_ID_LIST, SENSUBOT_TELEGRAM_ADMIN_ID_LIST, SENSUBOT_GCHAT_ADMIN_LIST: User ID (from slack, google chat and telegram) allowed to run anything

## [0.0.1] - 2020-02-03
### Added
- Slack Endpoint sensu_bot_url/sensubot/v1/slack
- Integration with Sensu API using token
- Add tests in usecase, appcontext, controller
- Add build script to control tag version and branch version
