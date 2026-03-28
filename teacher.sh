#!/bin/bash

interviewn=$(grep -H "licen" interviews/* | grep "\"" | cut -f1 -d ":" | rev | cut -f1 -d "-" | rev)

export interviewnum=$interviewn
echo "$interviewnum"

cat "interviews/interview-$interviewnum"

echo "$MAIN_SUSPECT"