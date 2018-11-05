import Data.Ratio


-- Rational Numbers

rationalSum ::  Int -> [Ratio Int]
rationalSum n = rationalNum 1 []
    where
        rationalNum :: Int -> [Ratio Int] -> [Ratio Int]
        rationalNum i xs
            | i >= n = xs
            | otherwise = rationalNum (i + 1) (xs ++ [i % (n - i)])
            
            
-- Lowest Terms Only

rationalSumLowest :: Int -> [Ratio Int]
rationalSumLowest n = rationalNum 1 []
    where
        rationalNum :: Int -> [Ratio Int] -> [Ratio Int]
        rationalNum i xs
            | i >= n = xs
            | gcd i n /= 1 = rationalNum (i + 1) xs
            | otherwise = rationalNum (i + 1) (xs ++ [i % (n - i)])



-- All Rational Numbers

rationals :: [Ratio Int]
rationals =  rt 1
    where
        rt :: Int -> [Ratio Int]
        rt y = (rationalSumLowest y) ++ (rt (y+1))
        
        
-- Input/Output

spl :: [Char] -> [Int] -> [Char] -> Int -> [Int]
spl str xs ys i
            | i >= length str && ys /= "" = xs ++ [read ys::Int]
            | i >= length str && ys == "" = xs
            | str !! i == '\n' && ys /= "" = spl str (xs ++ [read ys::Int]) "" (i+1)
            | otherwise = spl str xs (ys ++ [str !! i]) (i+1)
            

sumFile :: IO ()
sumFile = do
    file <- readFile "input.txt"
    putStrLn ( show (foldl (+) 0 (spl file [] "" 0)))
    
    
    