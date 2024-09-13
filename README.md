# FitQuest Workout Log

Fitquest is a simple workout log. It was designed to take away all of the fluff of exercising. Most workout apps are over complicated. Notebooks are hard to keep up with and I don't know about you but my handwriting is terrible. 

FitQuest is probably the simplest workout log you'll ever use. Set a goal, e.g, 100 Push-ups. That is 100 push-ups in a single set. Now do push-ups. Log how many you did using Fitquest.

FitQuest is as simple as it sounds.

## Usage

### Adding Exercises to Track

```sh
fitquest track exercise --name Push-ups --goal 100
 Id    Exercise        Goal    Record  Percentage
──────────────────────────────────────────────────────
 1     Push-ups        100     0       0%
```

You can also short hand the arguments.

```sh
fitquest -n Push-ups -g 100
 Id    Exercise        Goal    Record  Percentage
──────────────────────────────────────────────────────
 1     Push-ups        100     0       0%
```

Let's add some more. 

```sh
fitquest -n Squats -g 100
fitquest -n Pull-ups -g 10
```

### Viewing workout logs:

Display Log Summary 

```sh
fitquest
 Id    Exercise        Goal    Record  Percentage     
──────────────────────────────────────────────────────
 1     Push-ups        100     0       0%
 2     Squats          100     0       0%
 3     Pull-ups        10      0       0%
# ... 
```

Get session details for specific entry.

```sh
fitquest --id 1
Push-Ups:
  Goal: 100
  Record: 20 (2006-01-05)
  Percentage Complete: (Record / Goal) * 100
  
  Started: 2006-01-02
  Tracked for: 6 Days (today's date - first entry date)
  Sessions:
      Date: 2006-01-09 Count: 17
      Date: 2006-01-07 Count: 13
      Date: 2006-01-05 Count: 20
      Date: 2006-01-03 Count: 7
      Date: 2006-01-02 Count: 8
      Date: 2006-01-02 Count: 5
```



### Updating an Exercise or Goal

Using the `--id` flag will perform an update the exercises name and goal. 

```sh
fitquest --id 1 -g 50 -n "Sit ups"
```

### Adding Sessions

```sh 
fitquest --id 1 --reps 50
# OR
fitquest --id 1 -r 50
```