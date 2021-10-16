# Citrus
* 原始图集
  * https://github.com/26fe-cn/Citrus/releases/tag/v0.0.1

* 标签名  

```
橘子 citrus
梗 stem
屁股眼 citrus_embryo
腐烂 rot
```

python train.py --img 640 --epochs 300 --data ./data/coco128.yaml --cfg ./models/yolov5s.yaml --weights ./weights/yolov5s.pt --device 0 --batch-size 8