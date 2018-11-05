import Data.Maybe

-- Hailstone, Again

hailstone n
    | even n = div n 2
    | odd n = 3 * n + 1

hailseq :: Int -> [Int]    
hailseq n = hailnext [] n
    where
        hailnext a 1 = a ++ [1]
        hailnext a n = hailnext (a ++ [n]) (hailstone n)
        
hailseq' :: Int -> [Int]    
hailseq' n = takeWhile (/=1) (iterate hailstone n) ++ [1]


-- Joining Strings, Again

join :: [Char] -> [[Char]] -> [Char]
join a b = foldl (\x y-> if x == "" then x ++ y else x ++ a ++ y) "" b

-- Merge Sort

merge :: (Show a, Ord a) => [a] -> [a]  -> [a]
merge a b = mergeout [] a b where 
    mergeout c a b
        | a == [] && b == [] = c
        | b == [] = mergeout (c ++ [head a]) (drop 1 a) b
        | a == [] = mergeout (c ++ [head b]) a (drop 1 b)
        | (a !! 0 <= b !! 0)  = mergeout (c ++ [head a]) (drop 1 a) b
        | (a !! 0 > b !! 0) = mergeout (c ++ [head b]) a (drop 1 b)
        
xx=[1,2,3]
        
mergeSort :: (Show a, Ord a) => [a] -> [a]
mergeSort xs
    | length xs == 2 = merge [(xs !! 0)] [(xs !! 1)]
    | length xs > 2 = merge (mergeSort partA) (mergeSort partB)
    | otherwise = xs
        where
            (partA, partB)= splitAt ((length xs) `div` 2) xs
            
-- Searching? Maybe?

findElt :: (Show a, Ord a) => a -> [a]  -> Maybe Int
findElt y xs = fn 0
    where
        fn z = case z >= length xs of
                True -> Nothing
                False -> if (xs !! z) == y 
                            then Just z
                            else fn (z+1)
        
        
                 