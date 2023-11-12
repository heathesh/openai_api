# Open AI API request and response

This app allow you to make requests to the Open AI API using your Open AI API Key. It works as an application that runs in the command line / terminal on either Windows / Mac OS / Linux.

## Pre-requisites

To use this app, you must have the following:

1. Golang installed and configured on your operating system (see https://go.dev/doc/install).
2. An Open API account (see https://platform.openai.com/login?launch).
3. An Open API Key (see https://platform.openai.com/api-keys).
4. Optional: A markdown viewer app installed on your computer.

## Recommended usage

To use this app:

1. Rename the _settings.json.sample_ file to _settings.json_.
2. Edit the _settings.json_ file with your respective text editor and change the _apiKey_ to the API Key you got from the Open AI portal.
3. Run the app using the following commands:

```
go run .
```

4. Once the app runs, you will be asked for the prompt you wish to use, enter the prompt and the response will be generated and outputted to the screen as well as saved to the _output.md_ file.
5. If you have a markdown viewer installed, you can then use it to view the output.md file to see a more readable version of the Open AI API response.

## Example run

To see this in action, I've created a 33 second video showing it working: https://www.youtube.com/watch?v=gGDn_IqoyD0

## Future plans

I am hoping to update this app to allow people to chose which of the Open AI API endpoints to call, e.g. dall-e for Image Generation etc.