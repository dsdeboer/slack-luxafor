Behold!

To control the light you only need a valid release artifact from GitHub.

To also control the Slack status you need to provide a valid slack token. You can achieve this by
- Creating a new slack app on https://api.slack.com/apps 
- Hit the "Create new App" button
- Select "From Manifest" 
- Select a workspace you control, then "Next"
- Paste the below snippet into the text area, then "Next"
```json
{
    "display_information": {
        "name": "Luxafor Integration",
        "description": "Integrates Slack status with Luxafor",
        "background_color": "#418541",
        "long_description": "Integrates Slack status with Luxafor. For more details and requirements see also https://luxafor.com/ to get your own device and synchronize the status of your slack account with the device"
    },
    "oauth_config": {
        "scopes": {
            "user": [
                "users.profile:write",
                "users:write"
            ]
        }
    },
    "settings": {
        "org_deploy_enabled": false,
        "socket_mode_enabled": false,
        "token_rotation_enabled": false
    }
}
```
- Review the information and hit "Create"
- Then navigate to "Install app" and install the app to your workspace
- Then hit "Allow" to grant the app the necessary permissions
- And finally copy the "OAuth Token" and set it as an environment variable => `LUXAFOR_SLACK_API_TOKEN`
