import Control.Monad
import Data.Time.Calendar
import Data.Time.Calendar.OrdinalDate
import Data.List

-- Merging

merge :: (Show a, Ord a) => [a] -> [a]  -> [a]
merge a b = mergeout [] a b where 
    mergeout c a b
        | a == [] && b == [] = c
        | b == [] = mergeout (c ++ [head a]) (drop 1 a) b
        | a == [] = mergeout (c ++ [head b]) a (drop 1 b)
        | (a !! 0 <= b !! 0)  = mergeout (c ++ [head a]) (drop 1 a) b
        | (a !! 0 > b !! 0) = mergeout (c ++ [head b]) a (drop 1 b)
        
-- Tail Recursive Hailstone
hailstone :: Int -> Int
hailstone n
    | even n = div n 2
    | odd n = 3 * n + 1

hailLen :: Int -> Int
hailLen n = hailTail 0 n
  where
    hailTail a 1 = a
    hailTail a b = hailTail (a + 1) (hailstone b) 
    
-- Factorial

fact :: Int -> Int
fact n = factI 1 n
    where
        factI a 1 = a
        factI a 0 = a
        factI a b = factI (a * b) (b-1)
        
fact' :: Int -> Int
fact' n = foldl (*) 1 [1..n]


-- Haskell Library and Dates

daysInYear :: Integer -> [Day]
daysInYear y = [jan1 .. dec31]
    where   jan1 = fromGregorian y 1 1
            dec31 = fromGregorian y 12 31


isFriday :: Day -> Bool
isFriday n
    | snd (mondayStartWeek n) == 5 = True
    | otherwise = False
    
divisors :: Int -> [Int]
divisors n = [ i | i <- [2..(n `div` 2)], n `mod` i == 0]    

getDay (y,m,d) = d

isPrimeDay :: Day -> Bool
isPrimeDay n
    |   divisors (getDay (toGregorian n)) == [] = True
    |   otherwise = False
    
primeFridays :: Integer -> [Day]
primeFridays n = intersect (filter isPrimeDay (daysInYear n)) (filter isFriday (daysInYear n))