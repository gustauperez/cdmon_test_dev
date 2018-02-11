 #!/usr/bin/env bash

function sequence(){
	for i in $(seq 1 100); do
		local output=""
		 [[ $((${i} % 3)) == 0 ]] && output=${output}CD
		 [[ $((${i} % 5)) == 0 ]] && output=${output}mon
		 echo ${output:-${i}}
	done
}

sequence
