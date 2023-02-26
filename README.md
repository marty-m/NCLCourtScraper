<h1>NCL Court Scraper</h1>
<h2>Description</h2>
<p>An SMS alert service notifying when there are open booking slots for Basketball courts at the Newcastle University Sport & Fitness Center.</p>

<h2>Setup</h2>
(The application is written in Golang, so please make sure you have the latest version installed from [here](https://go.dev/). 
It also uses a headless Chrome browser for web-scraping, so please make sure you have a Chromium-based browser installed.)
<br><br>
Before building an executable make sure you fill out Newcastle University login details in both ``court5.go`` and 
``court8.go`` as well as your Twilio API credentials in ``sms.go``.

To build an executable run the following in the project root directory:<br>
```go
go build ./main/
```
(see examples on how to build executables for other platforms [here](https://opensource.com/article/21/1/go-cross-compiling))
<h2>Usage</h2>
Once the build is finished, run the executable:

```go
./my-executable
```

The output should look like this:

```text
---------2023-01-01  10:00:00---------
Starting alert service...
Searching available slots for Court 5...
Searching available slots for Court 8...
```

The scraper will check for open slots every two hours (this is modifiable in 
``startTracker.go``). 
<br>(**Note**: You will only be notified of an open booking slot once. If
the app has notified you about an open court slot in a previous message it will not include it in any future alerts.)
<h2>License</h2>

Copyright 2023 Martynas Miliauskas<br><br>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


