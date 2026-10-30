# golang interview (~20m)

## prompt

This program summarizes a list of orders by customer. Invalid orders have an empty customer name or a non-positive amount and should be ignored. If an invalid order is detected we should continue processing. 

The program produces incorrect summaries for some inputs. Find and fix the bug. Add a test or explain how you verified the fix.

## expected Behavior

```bash
❯ go run main.go
warning: invalid order for customer ""
Alice: 15
Bob: 7
```
