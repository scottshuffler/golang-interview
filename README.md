# golang interview (~20m)

## prompt

This program summarizes a list of orders by customer. Invalid orders have an empty customer name or a non-positive amount and should be ignored.

The program produces incorrect summaries for some inputs. Find and fix the bug. Add a test or explain how you verified the fix.

## expected Behavior

It should:

- add amounts for each customer
- ignore invalid orders with no Customer name
- return an error if one or more invalid orders are found
- continue processing valid orders after an invalid order
- print:

```text
warning: invalid order for customer ""
Alice: 15
Bob: 7
```
