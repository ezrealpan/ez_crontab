
#!/bin/bash

nohup ./process > promanager.log &

echo $! >> promanager.pid

