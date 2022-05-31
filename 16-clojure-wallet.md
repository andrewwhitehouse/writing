**Domain-Driven Design in Clojure**

and possibly Rust 

for Fun and Profit.

I took a tech test a while ago. I didn't get the role because I haven't been doing the specific things they wanted me to do. Afterwards I thought that it was harder work that it needed to be and the framework (Java/Spring) feels rather dogmatic to me.

As you may have noticed there's some tension between my maintaining the skills I need to do my day job (supporting a family as a fairly well paid developer) and the things that indulge my playful side.

I tried the Advent of Code problems. They're OK, but I don't think developers are looking to those as a way of gaining job-relevant skills. (Maybe they are, but there's the competitive coding side which is somewhat of a distraction.)

I use domain-driven design in my day job. So it seems logical to use that. 

I also thought how I can frame this project. Blockchain is an area I see growing and I want to get up to speed with that ... aiming for where people are heading.

And I thought: Clojure and Rust are the two [most loved](https://insights.stackoverflow.com/survey/2021#technology-most-loved-dreaded-and-wanted) languages on Stack Overflow. I know one, and I'm learning the other.

Perhaps it's a centralised vs decentralised theme.

I considered a word guessing game like Wordle: centralised and decentralised.

Or I could look at the blockchain applications people have built and build a server-based version first.

For now, I'm going to focus on a wallet implementation. I might do this in Clojure and Spring Boot to show those Java people (and I have a foot in that camp) what a great language and ecosystem Clojure is.

---

Create a sample app.

`lein new compojure wallet`

With the default Compojure app, you start it with `lein ring server`. I want to run it as a main function, so I'm going to add [Jetty](https://www.eclipse.org/jetty/).

Project dependencies live in `project.clj` so update this to include the Jetty adapter.

```
(defproject wallet "0.1.0-SNAPSHOT"
  :description "FIXME: write description"
  :url "http://example.com/FIXME"
  :min-lein-version "2.0.0"
  :dependencies [[org.clojure/clojure "1.10.0"]
                 [compojure "1.6.1"]
                 [ring/ring-defaults "0.3.2"]
                 [ring/ring-jetty-adapter "1.9.5"]]
  :plugins [[lein-ring "0.12.5"]]
  :ring {:handler wallet.handler/app}
  :main wallet.core
  :profiles
  {:dev {:dependencies [[javax.servlet/servlet-api "2.5"]
                        [ring/ring-mock "0.3.2"]]}})
```
        
Create a `core.clj` which is where our main usually lives:

```
(ns wallet.core
  (:require [wallet.handler :refer [app]]
            [ring.adapter.jetty :as jetty]))

(defn -main
  [& args]
  (let [port (or (System/getenv "PORT") 3000)]
    (log/info "Listening on port" port)
    (jetty/run-jetty app
                     {:port port
                      :join? true})))
```

Note that I added code which allows you to set the port on which the server listens.

```
$ PORT=8000 lein run
2022-05-31 19:29:22.948:INFO::main: Logging initialized @3612ms to org.eclipse.jetty.util.log.StdErrLog
Listening on port 8000
2022-05-31 19:29:23.036:INFO:oejs.Server:main: jetty-9.4.44.v20210927; built: 2021-09-27T23:02:44.612Z; git: 8da83308eeca865e495e53ef315a249d63ba9332; jvm 11.0.7+10-LTS
2022-05-31 19:29:23.102:INFO:oejs.AbstractConnector:main: Started ServerConnector@55d776ac{HTTP/1.1, (http/1.1)}{0.0.0.0:8000}
2022-05-31 19:29:23.103:INFO:oejs.Server:main: Started @3767ms
```

This is the existing request handler

```
(ns wallet.handler
  (:require [compojure.core :refer :all]
            [compojure.route :as route]
            [ring.middleware.defaults :refer [wrap-defaults site-defaults]]))

(defroutes app-routes
  (GET "/" [] "Hello World")
  (route/not-found "Not Found"))

(def app
  (wrap-defaults app-routes site-defaults))
```

Let's use Heroku to deploy it for now. [Instructions](https://devcenter.heroku.com/articles/getting-started-with-clojure?singlepage=true).


Follow the instructions to install Heroku and login.

Create an app on Heroku

`$ heroku create`

```
$ heroku create
Creating app... done, ⬢ protected-reef-78839
https://protected-reef-78839.herokuapp.com/ | https://git.heroku.com/protected-reef-78839.git
$ 
```

```
$ git remote -v
heroku	https://git.heroku.com/protected-reef-78839.git (fetch)
heroku	https://git.heroku.com/protected-reef-78839.git (push)
. . .
```

```
$ git push heroku main
Enumerating objects: 22, done.
Counting objects: 100% (22/22), done.
Delta compression using up to 8 threads
Compressing objects: 100% (16/16), done.
Writing objects: 100% (22/22), 2.54 KiB | 1.27 MiB/s, done.
Total 22 (delta 3), reused 0 (delta 0)
remote: Compressing source files... done.
remote: Building source:
remote: 
remote: -----> Building on the Heroku-20 stack

. . .

remote:        Retrieving ring/ring-servlet/1.9.5/ring-servlet-1.9.5.jar from clojars
remote:        Compiling wallet.core
remote:        2022-05-31 18:35:47.001:INFO::main: Logging initialized @5712ms to org.eclipse.jetty.util.log.StdErrLog
remote:        Compiling wallet.handler
remote: -----> Discovering process types
remote:        Procfile declares types     -> (none)
remote:        Default types for buildpack -> web
remote: 
remote: -----> Compressing...
remote:        Done: 94M
remote: -----> Launching...
remote:        Released v3
remote:        https://protected-reef-78839.herokuapp.com/ deployed to Heroku
remote: 
remote: Verifying deploy... done.
To https://git.heroku.com/protected-reef-78839.git
 * [new branch]      main -> main
$
```

If you browse to the page, you should see "Hello World".

and you would see "Hello World".

But it's not very impressive yet. Let's make the message more relatable.

When I run locally I don't always want to specify the port, and my logic wasn't quite right for when I run it locally.

```
(defn -main
  [& args]
  (let [env-port (System/getenv "PORT")
        selected-port (if env-port (Integer/parseInt env-port) 3000)]
    (println "Listening on port" selected-port)
    (jetty/run-jetty app
                     {:port selected-port
                      :join? true})))
```

And change the handler.

```
(defn wallet-markup []
  (let [url "https://www.fjallraven.com/496026/globalassets/catalogs/fjallraven/f7/f773/f77307/f249/7323450091569_ss18_a_oevik_wallet_21.jpg?width=120&height=120&mode=BoxPad&bgcolor=fff&quality=80"]
    (str
      "<img src=\"" url "\" alt=\"This is not a wallet\">"
      "<p>This is not a wallet</p>")))

(defroutes app-routes
  (GET "/" [] (wallet-markup))
  (route/not-found "Not Found"))
```

I tested it locally.

Let's push to Heroku.  And browse to the [application page](https://protected-reef-78839.herokuapp.com/). Hat tip to René Magritte.

@Beaver