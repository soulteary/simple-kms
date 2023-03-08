# Simple KMS

![](./screenshots/homepage.png)

## Usage

1. Run the `skms` with or without the `ACCESS_KEY` environment variable, then the server will listening at default port `8090`.

```
# ./skms
```

2. Run the `skms` with cli parameter `--generate` with the `ACCESS_KEY` environment variable, you will get the `ACCESS_KEY` and `SECRET_KEY` pairs.

```bash
# ACCESS_KEY=681cddf2-35b5-4d99-a8f0-6862fa79098e ./skms --generate
ACCESS_KEY=681cddf2-35b5-4d99-a8f0-6862fa79098e
SECRET_KEY=ddeed7d2-ee42-5728-a06b-b70ee6f7d703
```

3. Reference `example` to complete your project, and run your project with the `ACCESS_KEY` and `SECRET_KEY`, if the key pairs does not match, the program will automatically exit, or you can customize other logic as you like.
