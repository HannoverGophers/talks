for i in {1..5}; do
  make invoke-with-logs | tee /tmp/invoke_warm_$i.log
  grep -E "Duration|Billed Duration" /tmp/invoke_warm_$i.log
done