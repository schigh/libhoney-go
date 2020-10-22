# Honeycomb Code Exercise
Steve High, 2020-10-22

---
This is not intended to be a package in the SDK, obviously :)



Due to time constraints, this implementation is more of a "just get it working" approach.  Given the time and a more thorough grokking of `libhoney`, here's what I would have done to more thoroughly integrate Markers into the SDK:

1. `libhoney` is set up to handle events and only events, so the entire stack would need to be retrofitted to accomodate markers.
2. I implemented a `DeleteMarker` func, but it didnt make sanse to attache it as a method to `Marker`.  I think if there was a `marker` package, that call would "look" a lot better, e.g. `marker.Delete(id)` vs `libhoney.DeleteMarker(id)`
3. I really didnt add many (any?) comments, which is not representative of my code.

