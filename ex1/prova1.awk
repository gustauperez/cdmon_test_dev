BEGIN {i = 1; while (i <= 100) { output=""; if(i%3==0){output=output"CD"; } if(i%5==0){output=output"mon";} if(output ~ /^$/) { output=i;}print output; ++i; } }
