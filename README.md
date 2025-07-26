# random song bot

## Deployment
1. create `.env` to set the following environment variables
   - `CLIENT_ID`: spotify api client id
   - `CLIENT_SECRET`: spotify api client secret
   - `SPOTIFY_PLAYLIST_ID`: spotify playlist id
   - `DISCORD_PUBLIC_KEY`: discord public key
   - `DISCORD_BOT_TOKEN`: discord bot token
   - `PORT`: port to listen requests (optional)

2. Deploy app where request is reachable from the internet

   The easiest way to use [ngrok](https://ngrok.com/).
   Run the application with `go run .` with ngrok tunnelling with
   ```
   ngrok http http://localhost:<port>
   ```

3. Register the interaction endpoint to [discord application settings](https://discord.com/developers/applications)

   Note that endpoint is not root but `<root>/discord/callback`.
   Press save changes button to check the endpoint reachability from discord.

   If you are using Ngrok, the endpoint will appear in your terminal after running above command.
   Set `https://<random>.ngrok-free.app/discord/callback` as the interaction endpoint URL.