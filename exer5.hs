-- Built-In Functions

myIterate :: (a->a) -> a -> [a]
myIterate fn n = [n] ++ myIterate fn (fn n)

myTakeWhile :: (a-> Bool) -> [a] -> [a]
myTakeWhile fn ys = tw [] ys
    where
        tw xs' ys'
            | length ys' == 0 = xs'
            | fn (head ys') == True = tw (xs' ++ [head ys'])  (tail ys')
            | otherwise = xs'
    
-- Pascal's Triangle

pascal :: Int -> [Int]
pascal 0 = [1]
pascal 1 = [1, 1]
pascal n = pc (pascal (n - 1))
    where
        pcc xs' n' ys'
            | n' > (length xs') - 2 = ys'
            | otherwise = pcc xs' (n' + 1) ys' ++ [(xs' !! n') + (xs' !! (n'+1))]
        pc xs = [1] ++ (pcc xs 0 []) ++ [1]
        
-- Pointfree Addition

addPair :: (Num a) => a -> a -> a
addPair = \j -> \k -> j + k


-- Pointfree Filtering

withoutZeros :: (Num a, Eq a) => [a] -> [a]
withoutZeros = \xs -> wz xs [] 0
    where 
        wz xs' ys' i'
            | i' >= length xs' = ys'
            | xs' !! i' /= 0 = wz xs' (ys' ++ [xs' !! i']) (i' + 1)
            | otherwise = wz xs' ys' (i' + 1)
        
-- Exploring Fibonacci


fib :: Int -> Int
fib 0 = 0
fib 1 = 1
fib n = (fib $ n - 1)  +  (fib $ n - 2)

fibs :: [Int]
fibs = map fib [0..]

-- Something Else

things :: [Integer]
things = 0 : 1 : zipWith (+) things (tail things)
