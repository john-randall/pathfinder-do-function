
# Pathfinder-DO-Function

An A* Pathfinding implementation written in GO and set up for [DigitalOcean Functions](https://www.digitalocean.com/blog/introducing-digitalocean-functions-serverless-computing).

## API Reference

#### Calculate Path with A*

```http
  POST /sample/pathfinder
```

| Parameter | Type     | Description                | Example |
| :-------- | :------- | :------------------------- | :-------- |
| `start` | `int []` | **Required**. Starting position | `[0, 0]` |
| `end` | `int []` | **Required**. End Position  | `[2, 2]`|
| `grid` | `int [][]` | **Required**. Grid Representation where 0 represents walkable and 1 represents an obstacle. | `[[0,0,0],[0,0,0],[0,0,0]]` |

#### Example body
```json
{
    "start": [0, 0],
    "end": [2, 0],
    "grid": [
        [0, 0, 0],
        [1, 1, 0],
        [0, 0, 0]
    ]
}
```

#### Response
```json
{
  "statusCode": 200
  "body": {
    "path": [
      [0,0],
      [0,1],
      [0,2],
      [1,2],
      [2,2],
      [2,1],
      [2,0]
    ]
  }
}

```
