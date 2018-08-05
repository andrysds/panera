# Panera

Bukalapak PNR Squad Assistant

## Commands

### Public Commands

- /standup
- /standup_list
- /standup_skip

### Master Commands

- /init
- /init_standup
- /init_standup_list
- /standup_new_day

## Scheduler

[https://github.com/andrysds/panera-scheduler]

## Contributing

Send a pull request from your fork branch. [How?](https://help.github.com/articles/creating-a-pull-request-from-a-fork)

## Requirements

- Go
- Redis

## How to run locally

- Fork and clone the repository

- Set `.env`, install vendor and run

  ```sh
  cp env.sample .env
  make dep
  make run
  ```

- Access the commands via curl, postman, or web browser

  ```sh
  curl -u panera:panera http://localhost:9542/init
  ```
