// Imagine a system where each DNS resolution is incurring a very high latency.
// To optimize the performance, one way is to store results of some  of the most
//  queried domain For eg: If we query google.com 10 times in past 10 minutes,
// we can store it locally so that we don’t need to issue a DNS query again and again.

// Design such system in golang