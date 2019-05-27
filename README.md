# Panera

Bukalapak PNR Squad Assistant

## Commands

### Public Commands

- /birthdays
- /standup
- /standup_list
- /standup_skip

### Master Commands

- /standup_new_day

## Contributing

Send a pull request from your fork branch. [How?](https://help.github.com/articles/creating-a-pull-request-from-a-fork)

## Requirements

- Go
- MongoDB
- Redis

## How to run locally

- Fork and clone the repository

- Set `.env`, install vendor and run

  ```sh
  cp env.sample .env
  make dep
  make run
  ```

- Access via web browser

  http://localhost:9542/
