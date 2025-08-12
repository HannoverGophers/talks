for i in {1..5}; do
  make update-configuration >/dev/null
  make invoke | tee /tmp/invoke_$i.log
  # grep -E "Duration: .*Init Duration" >> benchmark.log
  grep -E "Duration|Billed Duration" /tmp/invoke_warm_$i.log
done