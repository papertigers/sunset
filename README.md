# sunset
Use the Dark Sky API to return the current day's Sunset Time.

## Usage

```
Usage of ./sunset:
  -epoch
    	Return SunsetTime in epoch (default true)
  -latitude string
    	latitude (default "42.8864")
  -longitude string
    	longitude (default "-78.8784")
```

### Example

```
DS_APIKEY="XXX" ./sunset
1503101697

DS_APIKEY="XXX" ./sunset -epoch=false
2017-08-18 17:14:57 -0700 PDT
```
