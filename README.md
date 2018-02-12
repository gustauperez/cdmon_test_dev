# CDmon test

This repo contains the code for the CDmon test. It's been split in two different directories, each one contains the code and stuff needed.

## Exercise 1

This one a simple test in which we have to code a simple program that has to follow this specification:

- Print all the integer numbers from 1 to 100
- If the number can be divided by 3, modulo 0, instead of the number the program has to print 'CD'
- On the other hand, if the number can be divided by 5, modulo 0, instead of the number the program has to print 'mon'
- If the number can be divided by both, the string 'CDmon' has to be printed

Three different scripts are provided. You can find them under the *ex1* directory. 

- The bash script can be run with *./prova1.sh*. It uses parameter substitution to print the default value (when the integer can't be divided neither by 3 nor by 5)
- The second one uses awk. It can be run with *awk -f prova1.awk*
- The third one uses golang. It has to be compiled and run: *go build prova.go && ./prova*

## Exercise 2

For this one you'd need a personal Digital Ocean key. Remember that you need to install docker in your machine (your mileage may vary, for Linux check your package manager, for Windows and MacOs, go to the official site and download them). After that, open a command line and run the following commands (I’m assuming a UNIX machine, for Windows probably you’d need to use back slashes instead):

```
# docker-machine create --driver digitalocean \ --digitalocean-access-token ${DIGITAL_OCEAN_TOKEN} ${DROPLET_NAME} [1]
# eval $(docker-machine env ${DROPLET_NAME})
# cd ${GUS_GITHUB_REPO}/ex1
# docker-compose build --up -d
```
Please replace **${DIGITAL_OCEAN_TOKEN}** with your personal token generated in your Digital Ocean account and **${DROPLET_NAME}** with the name of your droplet. **${GUS_GITHUB_REPO}** is the directory where you cloned the repo.

This will connect to your droplet, build the latest image of the official docker container from Apache (which uses the latest version of 2.4, tagged as latest) and run it in the droplet.

Here you can find a brief specificaction of the API:

- *GET* */hostings/* This will list all the hostings in the 'database'.
- *GET* */hostings/{id}* This will fetch a single hosting.

[1] eval $(docker-machine env ${DROPLET_NAME)}
