#!/bin/bash

INPUT_FILE=$1
OUTPUT_FILE="results_$(date +%Y%m%d_%H%M%S).csv"

echo "IOC,Risk Score,Verdict" > $OUTPUT_FILE

while IFS= read -r ioc; do
    if [ -n "$ioc" ]; then
        echo "Analyzing: $ioc"
        
        response=$(curl -s -X POST http://localhost:5000/api/analyze \
            -H "Content-Type: application/json" \
            -d "{\"text\": \"$ioc\"}")
        
        risk=$(echo $response | jq -r '.risk_score // "N/A"')
        verdict=$(echo $response | jq -r '.verdict // "Unknown"')
        
        echo "$ioc,$risk,$verdict" >> $OUTPUT_FILE
    fi
done < "$INPUT_FILE"

echo "Done! Results saved to $OUTPUT_FILE"
