one-time file serving for caddy
---
*Serve files exactly once*

Sometimes it's convenient for me to host a file on my website to send it to a friend, or another computer, or to [send a bit of text without it ending up in a chat log](https://onetimesecret.com/). But I don't want to leave the file around for anyone to see.

To scratch my own itch here, I made a simple [Caddy](https://caddyserver.com/) plugin. It serves files in a directory of your choice, and deletes each file after serving. That's it.

Now when I want to send a file to someone, I can upload it to my webserver and send my friend a link: https://www.example.com/one-time-secret/my-temporary-file
