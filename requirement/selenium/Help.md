Run

    Pull the image and run the container

     docker pull elgalu/selenium #upgrades to latest if a newer version is available

     docker run -d --name=grid -p 4444:24444 -p 5900:25900 \
         -e TZ="US/Pacific" -v /dev/shm:/dev/shm --privileged elgalu/selenium

    Wait until the grid starts properly before starting the tests (Optional but recommended)

     docker exec grid wait_all_done 30s
     # Or if docker exec is not available (eg. circleci)
     wget --retry-connrefused --no-check-certificate -T 30  http://localhost:4444/grid/console -O /dev/null

