"macro go"

$pseq=systemdir+'/psglib/'+seqfil
exists($pseq,'file'):$e
if ($e) then
  if ($# = 1) then
    go2($0,$1)
  else
    go2($0)
  endif
  return
endif

on('execprep'):$ex
if ($ex < -0.5) then
  " Run generalized go macro 'usergo' if it exists "
  exists('usergo','maclib'):$ex
  if ($ex)  then  usergo  endif
  " Run sequence specific macro 'go_<seqfil>' if it exists "
  $macro = 'go_'+seqfil
  exists($macro,'maclib'):$ex
  if ($ex)  then  {$macro}  endif
endif
il='n'

// Parse input
$args = $0
$i = 0
while ($i<$#) do
   $i = $i + 1
   $args = $args+'\',\''+${$i}
endwhile
$args = 'Acq(\'' + $args + '\')'

" write('line3',$args) "


"
$at_delay=''
format(at*1000+1,1,0):$at_delay
$nt_value=''
format(nt,1,0):$nt_value
B12SCNControl('outblank '+$at_delay+' '+$nt_value+' 1000') 
"

$c1=0
$args = $args+':$c1'
exec($args):$e
if ($##>0.5) then
  return($c1)
endif

