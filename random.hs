import System.Random

randInts :: Int -> Int -> Int -> IO [Int]
randInts n minval maxval = do
    gen <- newStdGen
    return $ take n $ randomRs (minval, maxval) gen
    
histogram :: (Eq a, Ord a, Enum a)
histogram vals = bars
    where
    counts =    [length $ filter (==i) vals
                | i <- [(minimum vals)..(maximum vals)] ]
    bars = [take n $ repeat ]

