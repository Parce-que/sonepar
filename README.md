# sonepar

## How things should work
We need to install a (docker) runner. It will be used to run the docker compose and start everything. 
The runner will get a http call from the CI. Can be integrated with Gitlab CI. 
We can also add some webhook to the git repository so it's triggered every push (for example).

The docker compose, will start the selenium and godogs dockers and make them run together. We can choose through some varenv to run the test on different environment. Depending on the project, it might also be put in a docker and run alongside (during the dev phase for example)


